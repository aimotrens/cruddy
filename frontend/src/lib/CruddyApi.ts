import Config from "@/Config";

const ApiDefinition = {
  List: (path: string) => "list?path=" + path,
  Download: (path: string) => "download?path=" + path,
  Delete: (path: string) => "delete?path=" + path,

  GetFullUrl: (action: string) => Config.Api.Combine(action),
};

export default ApiDefinition;
