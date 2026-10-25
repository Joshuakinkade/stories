"use server"

import { redirect } from "next/navigation";
import { destroySession, loadSession } from "../session";
import { SaveSettingsState } from "../types";
import { revalidatePath } from "next/cache";

export async function saveSettings(prevState: unknown, formData: FormData): Promise<SaveSettingsState> {
  const request = new Request(`${process.env.API_BASE_URL}/user`, {
    method: "PATCH",
    body: formData,
  });

  const sess = await loadSession();

  request.headers.set("Authorization", `Bearer ${sess.authToken}`);

  const response = await fetch(request);
  console.log(response);
  if (response.status < 300) {
    revalidatePath("/settings");
    return null;
  }
  else if (response.status == 401) {
    console.log(await response.text());
    await destroySession();
    return redirect("/");
  } else {
    return {
      error: "Server error"
    }
  }
}
