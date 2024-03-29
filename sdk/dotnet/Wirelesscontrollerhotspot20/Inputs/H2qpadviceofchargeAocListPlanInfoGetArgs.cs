// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontrollerhotspot20.Inputs
{

    public sealed class H2qpadviceofchargeAocListPlanInfoGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Currency code.
        /// </summary>
        [Input("currency")]
        public Input<string>? Currency { get; set; }

        /// <summary>
        /// Info file.
        /// </summary>
        [Input("infoFile")]
        public Input<string>? InfoFile { get; set; }

        /// <summary>
        /// Languague code.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// Plan name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public H2qpadviceofchargeAocListPlanInfoGetArgs()
        {
        }
        public static new H2qpadviceofchargeAocListPlanInfoGetArgs Empty => new H2qpadviceofchargeAocListPlanInfoGetArgs();
    }
}
