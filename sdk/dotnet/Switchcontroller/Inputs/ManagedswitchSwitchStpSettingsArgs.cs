// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class ManagedswitchSwitchStpSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable STP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ManagedswitchSwitchStpSettingsArgs()
        {
        }
        public static new ManagedswitchSwitchStpSettingsArgs Empty => new ManagedswitchSwitchStpSettingsArgs();
    }
}
