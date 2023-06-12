// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Antivirus.Inputs
{

    public sealed class ProfileContentDisarmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable inserting a cover page into the disarmed document. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("coverPage")]
        public Input<string>? CoverPage { get; set; }

        /// <summary>
        /// Enable/disable only detect disarmable files, do not alter content. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("detectOnly")]
        public Input<string>? DetectOnly { get; set; }

        /// <summary>
        /// Action to be taken if CDR engine encounters an unrecoverable error. Valid values: `block`, `log-only`, `ignore`.
        /// </summary>
        [Input("errorAction")]
        public Input<string>? ErrorAction { get; set; }

        /// <summary>
        /// Enable/disable stripping of PowerPoint action events in Microsoft Office documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("officeAction")]
        public Input<string>? OfficeAction { get; set; }

        /// <summary>
        /// Enable/disable stripping of Dynamic Data Exchange events in Microsoft Office documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("officeDde")]
        public Input<string>? OfficeDde { get; set; }

        /// <summary>
        /// Enable/disable stripping of embedded objects in Microsoft Office documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("officeEmbed")]
        public Input<string>? OfficeEmbed { get; set; }

        /// <summary>
        /// Enable/disable stripping of hyperlinks in Microsoft Office documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("officeHylink")]
        public Input<string>? OfficeHylink { get; set; }

        /// <summary>
        /// Enable/disable stripping of linked objects in Microsoft Office documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("officeLinked")]
        public Input<string>? OfficeLinked { get; set; }

        /// <summary>
        /// Enable/disable stripping of macros in Microsoft Office documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("officeMacro")]
        public Input<string>? OfficeMacro { get; set; }

        /// <summary>
        /// Destination to send original file if active content is removed. Valid values: `fortisandbox`, `quarantine`, `discard`.
        /// </summary>
        [Input("originalFileDestination")]
        public Input<string>? OriginalFileDestination { get; set; }

        /// <summary>
        /// Enable/disable stripping of actions that submit data to other targets in PDF documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("pdfActForm")]
        public Input<string>? PdfActForm { get; set; }

        /// <summary>
        /// Enable/disable stripping of links to other PDFs in PDF documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("pdfActGotor")]
        public Input<string>? PdfActGotor { get; set; }

        /// <summary>
        /// Enable/disable stripping of actions that execute JavaScript code in PDF documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("pdfActJava")]
        public Input<string>? PdfActJava { get; set; }

        /// <summary>
        /// Enable/disable stripping of links to external applications in PDF documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("pdfActLaunch")]
        public Input<string>? PdfActLaunch { get; set; }

        /// <summary>
        /// Enable/disable stripping of embedded movies in PDF documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("pdfActMovie")]
        public Input<string>? PdfActMovie { get; set; }

        /// <summary>
        /// Enable/disable stripping of embedded sound files in PDF documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("pdfActSound")]
        public Input<string>? PdfActSound { get; set; }

        /// <summary>
        /// Enable/disable stripping of embedded files in PDF documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("pdfEmbedfile")]
        public Input<string>? PdfEmbedfile { get; set; }

        /// <summary>
        /// Enable/disable stripping of hyperlinks from PDF documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("pdfHyperlink")]
        public Input<string>? PdfHyperlink { get; set; }

        /// <summary>
        /// Enable/disable stripping of JavaScript code in PDF documents. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("pdfJavacode")]
        public Input<string>? PdfJavacode { get; set; }

        public ProfileContentDisarmArgs()
        {
        }
        public static new ProfileContentDisarmArgs Empty => new ProfileContentDisarmArgs();
    }
}
