// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Endpointcontrol.Inputs
{

    public sealed class ProfileOnNetAddrArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address object from available options.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ProfileOnNetAddrArgs()
        {
        }
        public static new ProfileOnNetAddrArgs Empty => new ProfileOnNetAddrArgs();
    }
}