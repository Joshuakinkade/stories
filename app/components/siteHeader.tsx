import NextLink from 'next/link';
import { Button, Popover } from '@heroui/react';
import LoginForm from './loginForm';
import { loadSession } from '@/app/lib/session';
import SessionMenu from './SessionMenu';

function LoginPopup() {
  return (<Popover>
    <Button variant="secondary">Login</Button>
    <Popover.Content className="max-w-64">
      <Popover.Dialog>
        <LoginForm/>
      </Popover.Dialog>
    </Popover.Content>
  </Popover>);
}


export default async function SiteHeader() {
  const session = await loadSession();
  return (<header className="border-b p-4 mb-4 flex items-center">
    <h1 className="flex-1 text-lg">
      <NextLink href="/" className="link underline-offset-2">
        WebJournal
      </NextLink>
    </h1>
    <div className="flex gap-4">
      {session.userId ? <SessionMenu userEmail={session.userEmail} /> : <LoginPopup />}
    </div>
  </header>);
}
