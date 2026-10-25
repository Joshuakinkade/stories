import { redirect } from "next/navigation";
import { loadSession } from "../lib/session";
import { getUser } from "../lib/services/users";
import { Pagination } from "@heroui/react";

const SampleBox = () => (<div className="pt-50 bg-gray-600"></div>);

export default async function DashboardPage() {
  const session = await loadSession();
  if (!session.authToken) {
    redirect("/");
  }

  const user = await getUser(session.authToken);

  return (user && (<div className="px-8">
    <header className="my-8">
      <h1 className="text-4xl">{user.displayName}'s Stories</h1>
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
