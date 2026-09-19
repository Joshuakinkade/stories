import SettingsForm from "@/components/settingsForm";
import { loadSession } from "../lib/session"
import { Surface } from "@heroui/react";


export default async function SettingsPage() {
  const session = await loadSession();

  // fetch user details here

  return <Surface className="w-240 mx-auto p-8 rounded-xl">
    <header className="mb-8">
      <h1 className="text-4xl">{session.userEmail} Settings</h1>
    </header>
    <div>
      <SettingsForm/>
    </div>
  </Surface>
}
