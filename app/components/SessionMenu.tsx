"use client"

import { logout } from "@/lib/actions/auth";
import { useSystemDarkMode } from "@/lib/hooks/useSystemDarkMode";
import { Button, Dropdown, Label } from "@heroui/react";

export default function SessionMenu({ userEmail }: { userEmail: string }) {
  // The popover is portaled to <body>, outside the editor's dark-mode wrapper,
  // so it needs the theme class applied directly.
  const isDark = useSystemDarkMode();

  return (<Dropdown>
    <Button variant="secondary">{userEmail}</Button>
    <Dropdown.Popover className={isDark ? "dark" : undefined}>
      <Dropdown.Menu>
        <Dropdown.Item href="/editor/settings">
          <Label>Settings</Label>
        </Dropdown.Item>
        <Dropdown.Item onAction={() => logout()}>
          <Label>Log Out</Label>
        </Dropdown.Item>
      </Dropdown.Menu>
    </Dropdown.Popover>
  </Dropdown>);
}
