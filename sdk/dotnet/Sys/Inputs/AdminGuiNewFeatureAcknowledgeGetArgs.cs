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

    public sealed class AdminGuiNewFeatureAcknowledgeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Select menu ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public AdminGuiNewFeatureAcknowledgeGetArgs()
        {
        }
        public static new AdminGuiNewFeatureAcknowledgeGetArgs Empty => new AdminGuiNewFeatureAcknowledgeGetArgs();
    }
}