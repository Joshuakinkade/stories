"use client";
import { Button, Modal } from "@heroui/react";

enum LibraryItemType {
  Album,
  Photo
};

interface LibraryItem {
  label: string;
  uuid: string;
  type: LibraryItemType;
  photo: string;
};

type BrowserGridProps = {
  items: LibraryItem[];
};

function BrowserGrid({items}: BrowserGridProps) {
  return <div className="grid grid-cols-8 gap-4 my-4">
    {items.map( item => <BrowserTile key={item.uuid} item={item} /> )}
  </div>
}

type BrowserTileProps = {
  item: LibraryItem
}

function BrowserTile({item}: BrowserTileProps) {
  switch (item.type) {
    case LibraryItemType.Album:
      return <div className="relative w-full aspect-3/2">
        <div className="absolute left-0 top-0 h-7 w-1/2 max-w-full truncate rounded-t-lg bg-default px-4 text-sm leading-7">{item.label}</div>
        <div className="absolute inset-x-0 bottom-0 top-6 rounded-lg rounded-tl-none bg-default p-4"></div>
      </div>
    case LibraryItemType.Photo:
      return <div className="w-full aspect-3/2 rounded-lg bg-default p-4"></div>
    default:
      throw new Error("Unkonown Library Item Type");
  }
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

const fillerItems: LibraryItem[] = [
  {label: "One", uuid: "uuid-one", type: LibraryItemType.Album, photo: ""},
  {label: "Two", uuid: "uuid-two", type: LibraryItemType.Album, photo: ""},
  {label: "Three", uuid: "uuid-three", type: LibraryItemType.Album, photo: ""},
  {label: "Four", uuid: "uuid-four", type: LibraryItemType.Photo, photo: ""},
  {label: "Five", uuid: "uuid-five", type: LibraryItemType.Photo, photo: ""},
];

export default function PhotoBrowser() {

  const handleFileDrop = (files) => {
    // upload files here. Actually, call the photo service to do it.
    console.log(files.length);
  }

  return <Modal>
    <Button>Photos</Button>
    <Modal.Backdrop>
      <Modal.Container size="cover">
        <Modal.Dialog>
          <Modal.Header>
            <div className="flex items-center">
              <Modal.Heading className="text-xl font-semibold flex-1">Your Photos</Modal.Heading>
              <div>
                <Button size="sm">Upload</Button>
              </div>
            </div>
          </Modal.Header>
          <Modal.Body className="flex flex-col">
            <div className="flex flex-1 flex-col">
            <DropZone onFileDrop={handleFileDrop}>
              <BrowserGrid items={fillerItems}/>
            </DropZone>
            </div>
            <div className="text-sm text-muted">
              102 items
            </div>
          </Modal.Body>
        </Modal.Dialog>
      </Modal.Container>
    </Modal.Backdrop>
  </Modal>
}
