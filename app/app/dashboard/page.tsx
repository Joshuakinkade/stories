import { redirect } from "next/navigation";
import { loadSession } from "../lib/session";
import { getUser } from "../lib/services/users";
import PhotoBrowser from "@/components/photoBrowser/photoBrowser";
import { Button, Pagination } from "@heroui/react";
import { createStory } from "../lib/actions/stories";

const SampleBox = () => (<div className="pt-50 bg-gray-600"></div>);

function CreateStoryButton() {
  return <form action={createStory}>
    <Button type="submit">Write Story</Button>
  </form>
}

export default async function DashboardPage() {
  const session = await loadSession();
  if (!session.authToken) {
    redirect("/");
  }

  const user = await getUser(session.authToken);

  return (user && (<div className="px-8">
    <header className="my-8 flex">
      <div className="flex-1">
      <h1 className="text-4xl">{user.displayName}'s Stories</h1>
      </div>
      <div className="flex items-center gap-2">
        <CreateStoryButton/>
        <PhotoBrowser />
      </div>
    </header>
    <main>
      <div className="grid grid-cols-6 gap-12 my-12">
        <SampleBox/>
        <SampleBox/>
        <SampleBox/>
        <SampleBox/>
        <SampleBox/>
        <SampleBox/>
        <SampleBox/>
        <SampleBox/>
      </div>
      <div>
        <Pagination className="justify-center">
          <Pagination.Summary>1 to 20 of 120</Pagination.Summary>
          <Pagination.Content>
            <Pagination.Item>
              <Pagination.Previous>
                <Pagination.PreviousIcon/>
                <span>Previous</span>
              </Pagination.Previous>
            </Pagination.Item>
            <Pagination.Item>
              <Pagination.Next>
                <span>Next</span>
                <Pagination.NextIcon/>
              </Pagination.Next>
            </Pagination.Item>
          </Pagination.Content>
        </Pagination>
      </div>
    </main>
  </div>));
}
