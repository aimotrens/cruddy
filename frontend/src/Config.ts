const apiDomain = "";
//const apiDomain = "http://localhost:4231";

const cfg = {
  Api: {
    Url: apiDomain + "/api",
    Combine: (path: string) => `${cfg.Api.Url}/${path}`,
  },
  Env: await (await fetch(apiDomain + "/config.json")).json(),
};

export default cfg;
