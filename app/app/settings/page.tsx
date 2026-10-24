import SettingsForm from "@/components/settingsForm";
import { loadSession } from "../lib/session"
import { Surface } from "@heroui/react";
import { getUser } from "../lib/services/users";


export default async function SettingsPage() {
  const session = await loadSession();

  // fetch user details here
  const user = await getUser();

  return <Surface className="w-240 mx-auto p-8 rounded-xl">
    <header className="mb-8">
      <h1 className="text-4xl">Settings</h1>
    </header>
    <div>
      {user && <SettingsForm user={user}/>}
    </div>
  </Surface>
}
