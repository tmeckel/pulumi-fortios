// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Logfortianalyzer.Inputs
{

    public sealed class OverridesettingSerialArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Serial Number.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public OverridesettingSerialArgs()
        {
        }
        public static new OverridesettingSerialArgs Empty => new OverridesettingSerialArgs();
    }
}
