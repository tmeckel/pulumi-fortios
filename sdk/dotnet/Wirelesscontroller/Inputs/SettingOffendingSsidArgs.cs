// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Inputs
{

    public sealed class SettingOffendingSsidArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Actions taken for detected offending SSID. Valid values: `log`, `suppress`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Define offending SSID pattern (case insensitive), eg: word, word*, *word, wo*rd.
        /// </summary>
        [Input("ssidPattern")]
        public Input<string>? SsidPattern { get; set; }

        public SettingOffendingSsidArgs()
        {
        }
        public static new SettingOffendingSsidArgs Empty => new SettingOffendingSsidArgs();
    }
}
