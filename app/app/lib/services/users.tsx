export interface User {
  id: string
  email: string
  displayName: string
  slug: string
}

const API_BASE_URL = process.env.API_BASE_URL || '';

export async function getUser(authToken: string): Promise<User | null> {
  if (!authToken) {
    return null;
  }

  const request = new Request(`${API_BASE_URL}/user-details`);
  request.headers.append("Authorization", `Bearer ${authToken}`)

  try {
    const response = await fetch(request);
    const user = await response.json();
    return user;
  } catch (e) {
    console.log(e);
  }

  return null;
}
