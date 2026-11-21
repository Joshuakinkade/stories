import SiteHeader from "@/components/siteHeader";
import SystemThemeScope from "@/components/systemThemeScope";

export default function EditorLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <SystemThemeScope className="min-h-screen pb-8">
      <SiteHeader />
      {children}
    </SystemThemeScope>
  );
}
