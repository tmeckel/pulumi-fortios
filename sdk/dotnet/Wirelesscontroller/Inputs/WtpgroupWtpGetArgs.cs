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

    public sealed class WtpgroupWtpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// WTP ID.
        /// </summary>
        [Input("wtpId")]
        public Input<string>? WtpId { get; set; }

        public WtpgroupWtpGetArgs()
        {
        }
        public static new WtpgroupWtpGetArgs Empty => new WtpgroupWtpGetArgs();
    }
}
