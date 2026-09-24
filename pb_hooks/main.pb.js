routerAdd("GET", "/api/healthz/js", (e) => {
  return e.json(200, {
    status: "ok",
    source: "js",
  })
})
