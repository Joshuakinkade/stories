"use client";

import { useEffect, useRef } from "react";

const QUERY = "(prefers-color-scheme: dark)";

// Runs during HTML parsing, before first paint, to avoid a light-theme flash.
const initScript = `document.getElementById("system-theme-scope")?.classList.toggle("dark", window.matchMedia("${QUERY}").matches);`;

/**
 * Applies HeroUI's `dark` class to its wrapper element whenever the system
 * prefers a dark color scheme, scoping dark mode to its children.
 */
export default function SystemThemeScope({
  children,
  className = "",
}: {
  children: React.ReactNode;
  className?: string;
}) {
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const mql = window.matchMedia(QUERY);
    const apply = () => ref.current?.classList.toggle("dark", mql.matches);
    apply();
    mql.addEventListener("change", apply);
    return () => mql.removeEventListener("change", apply);
  }, []);

  return (
    <div
      id="system-theme-scope"
      ref={ref}
      className={`bg-background text-foreground ${className}`}
      suppressHydrationWarning
    >
      <script dangerouslySetInnerHTML={{ __html: initScript }} />
      {children}
    </div>
  );
}
