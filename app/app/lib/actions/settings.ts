"use server"

import { redirect } from "next/navigation";
import { destroySession, loadSession } from "../session";
import { SaveSettingsState } from "../types";

export async function saveSettings(prevState: unknown, formData: FormData): Promise<SaveSettingsState> {
  const request = new Request(`${process.env.API_BASE_URL}/settings`, {
    method: "POST",
    body: formData,
  });

  const sess = await loadSession();

  request.headers.set("Authorization", `Bearer ${sess.authToken}`);

  const response = await fetch(request);
  if (response.status < 300) {
    return null;
  }
  else if (response.status == 401) {
    console.log(await response.text());
    await destroySession();
    return redirect("/")
  } else {
    return {
      error: "Server error"
    }
  }
}
