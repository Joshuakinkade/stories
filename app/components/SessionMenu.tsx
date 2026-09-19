"use client"

import { logout } from "@/app/lib/actions/auth";
import { Button, Dropdown, Label } from "@heroui/react";

export default function SessionMenu({ userEmail }: { userEmail: string }) {
  return (<Dropdown>
    <Button variant="secondary">{userEmail}</Button>
    <Dropdown.Popover>
      <Dropdown.Menu>
        <Dropdown.Item href="/settings">
          <Label>Settings</Label>
        </Dropdown.Item>
        <Dropdown.Item onAction={() => logout()}>
          <Label>Log Out</Label>
        </Dropdown.Item>
      </Dropdown.Menu>
    </Dropdown.Popover>
  </Dropdown>);
}
