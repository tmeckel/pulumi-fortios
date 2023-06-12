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

    public sealed class ProfileSmbGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Select the archive types to block. Valid values: `encrypted`, `corrupted`, `partiallycorrupted`, `multipart`, `nested`, `mailbomb`, `fileslimit`, `timeout`, `unhandled`.
        /// </summary>
        [Input("archiveBlock")]
        public Input<string>? ArchiveBlock { get; set; }

        /// <summary>
        /// Select the archive types to log. Valid values: `encrypted`, `corrupted`, `partiallycorrupted`, `multipart`, `nested`, `mailbomb`, `fileslimit`, `timeout`, `unhandled`.
        /// </summary>
        [Input("archiveLog")]
        public Input<string>? ArchiveLog { get; set; }

        /// <summary>
        /// Enable/disable the virus emulator. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("emulator")]
        public Input<string>? Emulator { get; set; }

        /// <summary>
        /// Enable/disable SMB AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`.
        /// </summary>
        [Input("outbreakPrevention")]
        public Input<string>? OutbreakPrevention { get; set; }

        public ProfileSmbGetArgs()
        {
        }
        public static new ProfileSmbGetArgs Empty => new ProfileSmbGetArgs();
    }
}
