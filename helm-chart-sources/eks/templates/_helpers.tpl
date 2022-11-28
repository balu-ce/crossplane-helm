{{/*
Expand the name of the chart.
*/}}

{{/*
Annotations for the eks cluster
*/}}
{{- define "eks.annotations" -}}
upjet.upbound.io/timeout: "2400"
{{- end }}

