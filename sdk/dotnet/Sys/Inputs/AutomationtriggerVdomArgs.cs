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

    public sealed class AutomationtriggerVdomArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Virtual domain name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AutomationtriggerVdomArgs()
        {
        }
        public static new AutomationtriggerVdomArgs Empty => new AutomationtriggerVdomArgs();
    }
}