// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class Policy6UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Names of individual users that can authenticate with this policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Policy6UserArgs()
        {
        }
        public static new Policy6UserArgs Empty => new Policy6UserArgs();
    }
}
