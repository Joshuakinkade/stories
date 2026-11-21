import { Story } from "@/lib/services/stories";
import { Surface } from "@heroui/react";

type StoryPageProps = {
  params: Promise<{ storyID: string }>
}

export default async function StoryPage({ params }: StoryPageProps) {
  const { storyID } = await params;

  const story: Story = {
    id: storyID,
    title: "Example Title",
    slug: "example-title",
    createdAt: new Date(),
    updatedAt: new Date(),
    elements: [],
  };

  return <div className="px-8 flex gap-4">
    <Surface className="w-64 p-4 rounded-xl">
      <section>
        <h2>Story Information</h2>
      </section>
      <section>
        <h2>Story Elements</h2>
      </section>
    </Surface>
    <Surface className="flex-1 p-4 rounded-xl">
      <header>
        <h1 className="text-4xl text-center mb-2">Editing: {story.title}</h1>
        <p className="text-center">Metadata goes here: Author, Tags, etc.</p>
      </header>
      <main>
        Canvas
      </main>
    </Surface>
  </div>
}
