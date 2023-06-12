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
    public sealed class GetBgpVrfResult
    {
        /// <summary>
        /// List of export route target. The structure of `export_rt` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBgpVrfExportRtResult> ExportRts;
        /// <summary>
        /// Import route map.
        /// </summary>
        public readonly string ImportRouteMap;
        /// <summary>
        /// List of import route target. The structure of `import_rt` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBgpVrfImportRtResult> ImportRts;
        /// <summary>
        /// Target VRF table. The structure of `leak_target` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBgpVrfLeakTargetResult> LeakTargets;
        /// <summary>
        /// Route Distinguisher: AA|AA:NN.
        /// </summary>
        public readonly string Rd;
        /// <summary>
        /// VRF role.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// Target VRF ID &lt;0 - 31&gt;.
        /// </summary>
        public readonly string Vrf;

        [OutputConstructor]
        private GetBgpVrfResult(
            ImmutableArray<Outputs.GetBgpVrfExportRtResult> exportRts,

            string importRouteMap,

            ImmutableArray<Outputs.GetBgpVrfImportRtResult> importRts,

            ImmutableArray<Outputs.GetBgpVrfLeakTargetResult> leakTargets,

            string rd,

            string role,

            string vrf)
        {
            ExportRts = exportRts;
            ImportRouteMap = importRouteMap;
            ImportRts = importRts;
            LeakTargets = leakTargets;
            Rd = rd;
            Role = role;
            Vrf = vrf;
        }
    }
}
