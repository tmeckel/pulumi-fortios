// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Inputs
{

    public sealed class IpamRuleDeviceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fortigate serial number or wildcard.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public IpamRuleDeviceGetArgs()
        {
        }
        public static new IpamRuleDeviceGetArgs Empty => new IpamRuleDeviceGetArgs();
    }
}
