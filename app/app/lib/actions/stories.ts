'use server'

import { redirect } from "next/navigation";
import { v7 as uuidV7 } from "uuid";

export async function createStory() {
  const uuid = uuidV7();
  redirect(`/stories/${uuid}`);
}
