'use server'

const API_BASE_URL = process.env.API_BASE_URL || '';

import { redirect } from "next/navigation";
import { startSession } from "../session";
import { destroySession } from "../session";
import { LoginState, SignupState } from "../types";

export async function signup(prevState: unknown, formData: FormData): Promise<SignupState> {
  const request = new Request(`${API_BASE_URL}/register`, {
    method: "POST",
    body: formData,
  });

  const response = await fetch(request);
  if (response.ok) {
    const data = await response.json();
    await startSession({
      authToken: data.user.authToken,
      userId: data.user.userId,
      userEmail: data.user.userEmail,
    });
    redirect("/dashboard");
  } else {
    const error = await response.json();
    console.error(response.statusText);
    return error;
  }
}

export async function login(prevState: unknown, formData: FormData): Promise<LoginState> {
  const request = new Request(`${API_BASE_URL}/login`, {
    method: "POST",
    body: formData,
  });
  const response = await fetch(request);
  console.log(response);

  const data = await response.json();

  if (response.status == 200) {
    await startSession({
      authToken: data.token,
      userId: data.user.id,
      userEmail: data.user.email,
    });
  } else {
    return {
      error: await data.error
    }
  }
  return {}
}

export async function logout() {
  await destroySession();
  redirect('/');
}
