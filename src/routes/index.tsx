import { createFileRoute } from '@tanstack/react-router'
import { HomeTemplate } from '@/features/home'

export const Route = createFileRoute('/')({
  component: RouteComponent,
})

function RouteComponent() {
  return <HomeTemplate />;
}
