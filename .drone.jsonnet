local Pipeline(name, image) = {
  kind: "pipeline",
  name: name,
  workspace: {
    base: "/go",
    path: "src/github.com/dhavlev/go-object-oriented"
  },
  steps: [
    {
      name: "test",
      image: image,
      commands: [
        "go test",
        "go build"
      ]
    }
  ]
};

[
  Pipeline("build", "golang")
]
