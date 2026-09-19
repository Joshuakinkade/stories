import { redirect } from "next/navigation";
import { destroySession } from "../lib/session";

export async function GET() {
  await destroySession();
  console.log("hello");
  return redirect("/");
}
