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

    public sealed class PolicyInternetService6SrcGroupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PolicyInternetService6SrcGroupGetArgs()
        {
        }
        public static new PolicyInternetService6SrcGroupGetArgs Empty => new PolicyInternetService6SrcGroupGetArgs();
    }
}