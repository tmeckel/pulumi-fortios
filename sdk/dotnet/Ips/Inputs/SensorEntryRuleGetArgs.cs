// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Ips.Inputs
{

    public sealed class SensorEntryRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rule IPS.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public SensorEntryRuleGetArgs()
        {
        }
        public static new SensorEntryRuleGetArgs Empty => new SensorEntryRuleGetArgs();
    }
}
