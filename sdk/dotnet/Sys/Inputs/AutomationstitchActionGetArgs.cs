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

    public sealed class AutomationstitchActionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action name.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Delay before execution (in seconds).
        /// </summary>
        [Input("delay")]
        public Input<int>? Delay { get; set; }

        /// <summary>
        /// Entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Required in action chain. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("required")]
        public Input<string>? Required { get; set; }

        public AutomationstitchActionGetArgs()
        {
        }
        public static new AutomationstitchActionGetArgs Empty => new AutomationstitchActionGetArgs();
    }
}
