// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extendercontroller.Inputs
{

    public sealed class ExtenderWanExtensionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FortiExtender interface name.
        /// </summary>
        [Input("modem1Extension")]
        public Input<string>? Modem1Extension { get; set; }

        /// <summary>
        /// FortiExtender interface name.
        /// </summary>
        [Input("modem2Extension")]
        public Input<string>? Modem2Extension { get; set; }

        public ExtenderWanExtensionGetArgs()
        {
        }
        public static new ExtenderWanExtensionGetArgs Empty => new ExtenderWanExtensionGetArgs();
    }
}
