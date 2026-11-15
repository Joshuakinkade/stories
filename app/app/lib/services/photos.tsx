export interface Photo {
  id: string,
  altText: string,
  createdAt: Date,
}

const API_BASE_URL = process.env.API_BASE_URL || '';

export function listPhotos(authToken: string, storyId: string): Promise<Photo | null> {
  if (!authToken) {
    return null;
  }

  const request = new Request(`${API_BASE_URL}/stories/${storyId}/photos`);
  request.headers.append("Authorization", `Bearer ${authToken}`);

  try {
    const response = await fetch(request);
    const user = await response.json();
    return user;
  } catch (e) {
    console.log(e);
  }

  return null;
}
