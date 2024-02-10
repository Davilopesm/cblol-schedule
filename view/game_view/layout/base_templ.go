// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><title>CBLOL 2024</title><body><h2>Game Schedule</h2><style>\n        @import url('https://fonts.googleapis.com/css2?family=Roboto:wght@400&display=swap');\n\n        h2 {\n          color: #333;\n          text-align: center;\n          margin-bottom: 20px;\n          font-family: 'Roboto', sans-serif;\n        }\n\n        table {\n          width: 100%;\n          border-collapse: collapse;\n          margin-top: 20px;\n          box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);\n        }\n\n        th, td {\n          border: 1px solid #ddd;\n          padding: 12px;\n          text-align: left;\n        }\n\n        th {\n          background-color: #f2f2f2;\n          color: #333;\n        }\n\n        .team-info {\n          font-size: 0.9em;\n          color: #555;\n        }\n      </style><table><thead><tr><th>Date</th><th>Time</th><th>Teams</th><th>Location</th><th>Results</th></tr></thead>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</table></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
