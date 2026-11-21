"use client";

import { Button, Popover } from "@heroui/react";
import { useSystemDarkMode } from "@/lib/hooks/useSystemDarkMode";
import LoginForm from "./loginForm";

export default function LoginPopup() {
  // The popover is portaled to <body>, outside the editor's dark-mode wrapper,
  // so it needs the theme class applied directly.
  const isDark = useSystemDarkMode();

  return (<Popover>
    <Button variant="secondary">Login</Button>
    <Popover.Content className={`max-w-64${isDark ? " dark" : ""}`}>
      <Popover.Dialog>
        <LoginForm/>
      </Popover.Dialog>
    </Popover.Content>
  </Popover>);
}
