// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Application.Inputs
{

    public sealed class NameMetadataGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Meta ID.
        /// </summary>
        [Input("metaid")]
        public Input<int>? Metaid { get; set; }

        /// <summary>
        /// Value ID.
        /// </summary>
        [Input("valueid")]
        public Input<int>? Valueid { get; set; }

        public NameMetadataGetArgs()
        {
        }
        public static new NameMetadataGetArgs Empty => new NameMetadataGetArgs();
    }
}
