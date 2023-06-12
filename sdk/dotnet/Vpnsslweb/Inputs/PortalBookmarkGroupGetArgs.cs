// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpnsslweb.Inputs
{

    public sealed class PortalBookmarkGroupGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("bookmarks")]
        private InputList<Inputs.PortalBookmarkGroupBookmarkGetArgs>? _bookmarks;

        /// <summary>
        /// Bookmark table. The structure of `bookmarks` block is documented below.
        /// </summary>
        public InputList<Inputs.PortalBookmarkGroupBookmarkGetArgs> Bookmarks
        {
            get => _bookmarks ?? (_bookmarks = new InputList<Inputs.PortalBookmarkGroupBookmarkGetArgs>());
            set => _bookmarks = value;
        }

        /// <summary>
        /// Bookmark group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PortalBookmarkGroupGetArgs()
        {
        }
        public static new PortalBookmarkGroupGetArgs Empty => new PortalBookmarkGroupGetArgs();
    }
}
