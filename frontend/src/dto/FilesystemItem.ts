export default interface FilesystemItem {
  filePath: string;
  name: string;
  isDir: boolean;
  size: number;
  changed: Date;
}
