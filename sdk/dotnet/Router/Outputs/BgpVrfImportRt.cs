// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class BgpVrfImportRt
    {
        /// <summary>
        /// Attribute: AA|AA:NN.
        /// </summary>
        public readonly string? RouteTarget;

        [OutputConstructor]
        private BgpVrfImportRt(string? routeTarget)
        {
            RouteTarget = routeTarget;
        }
    }
}
