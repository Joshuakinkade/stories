"use server"

import { loadSession } from "../session";
import { SaveSettingsState } from "../types";

export async function saveSettings(prevState: unknown, formData: FormData): Promise<SaveSettingsState> {
  const request = new Request(`${process.env.API_BASE_URL}/settings`, {
    method: "POST",
    body: formData,
  });

  const sess = await loadSession();

  request.headers.set("Authorization", `Bearer ${sess.authToken}`);

  const response = await fetch(request);
  console.log(response.status, await response.text());
  return null;
}
