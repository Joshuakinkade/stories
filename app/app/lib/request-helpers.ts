import { loadSession } from "./session";

export async function authenticatedFetch(request: Request): Promise<Response> {
  const session = await loadSession();
  request.headers.set("Authorization", `Bearer ${session.authToken}`);
  // get auth token and add it to request
  // if the token comes back expired use the refresh token
  const response = await fetch(request);
  // if (response.status == 401) {
  // }
  return response
}
