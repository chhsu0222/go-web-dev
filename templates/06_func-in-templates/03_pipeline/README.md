# Pipeline

A pipeline may be "chained" by separating a sequence of commands with pipeline characters '|'.
In a chained pipeline, the result of each command is passed as the last argument of the following command.
The output of the final command in the pipeline is the value of the pipeline.

[golang.org/pkg/text/template](https://golang.org/pkg/text/template/#hdr-Pipelines)


