# YAML Graph JSONSchema

The YAML Representation Graph spec can be encoded as JSON, allowing bidirectional encoding of arbitrary YAML documents in JSON. The [schema](yaml-graph.schema.json) includes `scalar`, `sequence` and `mapping` kinds with arbitrary tags. To support references there is a type for `alias` nodes with anchors.

* https://yaml.org/spec/1.2.2/#321-representation-graph
* [yaml.io Matrix discussion](https://matrix.to/#/!KNlExDLtvbTGAPmgsq:yaml.io/$LCmaDFM5twkTpWfZ97V8mt-zXcvyB8r4TEc8YqLnsCM)

```
$ go run .
root:
    kind: scalar
    tag: str
    value: Hello world!
```
