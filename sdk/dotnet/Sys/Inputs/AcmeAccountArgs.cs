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

    public sealed class AcmeAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account ca_url.
        /// </summary>
        [Input("caUrl")]
        public Input<string>? CaUrl { get; set; }

        /// <summary>
        /// Account email.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Account id.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Account Private Key.
        /// </summary>
        [Input("privatekey")]
        public Input<string>? Privatekey { get; set; }

        /// <summary>
        /// Account status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Account url.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public AcmeAccountArgs()
        {
        }
        public static new AcmeAccountArgs Empty => new AcmeAccountArgs();
    }
}
