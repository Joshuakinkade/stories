
type UserGalleryProps = {
  params: Promise<{ userSlug: string }>
}

export default async function UserGallery({ params }: UserGalleryProps) {
  const { userSlug } = await params;

  return <div>{userSlug}</div>
}
