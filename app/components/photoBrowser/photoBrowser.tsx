"use client";
import { Button, Popover } from "@heroui/react";
import { Photo } from "@/app/lib/services/photos.tsx"

type BrowserGridProps = {
  photos: Photo[];
};

function BrowserGrid({photos}: BrowserGridProps) {
  return <div className="grid grid-cols-4 gap-4 my-4">
    {photos.map( photo => <BrowserTile key={photo.id} photo={photo} /> )}
  </div>
}

type BrowserTileProps = {
  photo: Photo;
};

function BrowserTile({photo}: BrowserTileProps) {
  return <div className="w-full aspect-3/2 rounded-lg bg-default p-4"></div>
}

function DropZone({children, onFileDrop}) {
  const handleDrop = (evt) => {
    const files = evt.dataTransfer.files;
    if( files.length == 0) {
      return;
    }

    evt.preventDefault();

    onFileDrop(files);
  }

  return <div
    className="flex-1 w-full"
    onDragOver={(evt) => evt.preventDefault()}
    onDrop={handleDrop}
  >
    {children}
  </div>
}

const fillerPhotos: Photo[] = [
  {id: "uuid-one", altText: "One", createdAt: new Date()},
  {id: "uuid-two", altText: "Two", createdAt: new Date()},
  {id: "uuid-three", altText: "Three", createdAt: new Date()},
  {id: "uuid-four", altText: "Four", createdAt: new Date()},
  {id: "uuid-five", altText: "Five", createdAt: new Date()},
];

export default function PhotoBrowser() {

  const handleFileDrop = (files) => {
    // upload files here. Actually, call the photo service to do it.
    console.log(files.length);
  }

  return <Popover>
    <Button>Photos</Button>
    <Popover.Content className="w-[48rem] max-w-[90vw]">
      <Popover.Dialog>
        <div className="flex items-center">
          <Popover.Heading className="text-xl font-semibold flex-1">Your Photos</Popover.Heading>
          <div>
            <Button size="sm">Upload</Button>
          </div>
        </div>
        <DropZone onFileDrop={handleFileDrop}>
          <BrowserGrid photos={fillerPhotos}/>
        </DropZone>
        <div className="text-sm text-muted">
          102 items
        </div>
      </Popover.Dialog>
    </Popover.Content>
  </Popover>
}
