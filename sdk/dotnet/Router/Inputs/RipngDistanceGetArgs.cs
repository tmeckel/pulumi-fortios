// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class RipngDistanceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access list for route destination.
        /// </summary>
        [Input("accessList6")]
        public Input<string>? AccessList6 { get; set; }

        /// <summary>
        /// Distance (1 - 255).
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// Distance ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Distance prefix6.
        /// </summary>
        [Input("prefix6")]
        public Input<string>? Prefix6 { get; set; }

        public RipngDistanceGetArgs()
        {
        }
        public static new RipngDistanceGetArgs Empty => new RipngDistanceGetArgs();
    }
}
