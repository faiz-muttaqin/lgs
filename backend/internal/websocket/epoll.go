package websocket

import (
	"log"
	"net"
	"reflect"
	"sync"
	"syscall"

	"github.com/gobwas/ws/wsutil"
	"golang.org/x/sys/unix"
)

var WS_POOL *Epoll

type Epoll struct {
	Fd          int
	Connections map[int]net.Conn // map of file descriptor to connection
	Lock        *sync.RWMutex
}

func MkEpoll() (*Epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &Epoll{
		Fd:          fd,
		Lock:        &sync.RWMutex{},
		Connections: make(map[int]net.Conn),
	}, nil
}

func (e *Epoll) Add(conn net.Conn) error {
	// Extract file descriptor associated with the connection
	fd := websocketFD(conn)
	err := unix.EpollCtl(e.Fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})
	if err != nil {
		return err
	}
	e.Lock.Lock()
	defer e.Lock.Unlock()
	e.Connections[fd] = conn
	if len(e.Connections)%100 == 0 {
		log.Printf("Total number of connections: %v", len(e.Connections))
	}
	return nil
}

func (e *Epoll) Remove(conn net.Conn) error {
	fd := websocketFD(conn)
	err := unix.EpollCtl(e.Fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}
	e.Lock.Lock()
	defer e.Lock.Unlock()
	delete(e.Connections, fd)
	if len(e.Connections)%100 == 0 {
		log.Printf("Total number of connections: %v", len(e.Connections))
	}
	return nil
}

func (e *Epoll) Wait() ([]net.Conn, error) {
	events := make([]unix.EpollEvent, 100)
	n, err := unix.EpollWait(e.Fd, events, 100)
	if err != nil {
		return nil, err
	}
	e.Lock.RLock()
	defer e.Lock.RUnlock()
	var connections []net.Conn
	for i := 0; i < n; i++ {
		conn := e.Connections[int(events[i].Fd)]
		connections = append(connections, conn)
	}
	return connections, nil
}

func websocketFD(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}

func StartListeningWebsocket() {
	for {
		connections, err := WS_POOL.Wait()
		if err != nil {
			log.Printf("Failed to epoll wait %v", err)
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				break
			}
			if _, _, err := wsutil.ReadClientData(conn); err != nil {
				if err := WS_POOL.Remove(conn); err != nil {
					log.Printf("Failed to remove %v", err)
				}
				conn.Close()
			} else {
				// This is commented out since in demo usage, stdout is showing messages sent from > 1M connections at very high rate
				//log.Printf("msg: %s", string(msg))
			}
		}
	}
}
