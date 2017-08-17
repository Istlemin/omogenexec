workspace(name = "omogenexec")

git_repository(
		name   = "com_github_gflags_gflags",
		tag    = "v2.2.1",
		remote = "https://github.com/gflags/gflags.git",
		)

bind(
		name   = "gflags",
		actual = "@com_github_gflags_gflags//:gflags",
		)

new_http_archive(
		name = "gtest",
		url = "https://github.com/google/googletest/archive/release-1.7.0.zip",
		sha256 = "b58cb7547a28b2c718d1e38aee18a3659c9e3ff52440297e965f5edffe34b6d0",
		build_file = "gtest.BUILD",
		strip_prefix = "googletest-release-1.7.0",
		)

git_repository(
		name = "org_pubref_rules_protobuf",
		remote = "https://github.com/pubref/rules_protobuf",
		tag = "v0.7.2",
		)

load("@org_pubref_rules_protobuf//cpp:rules.bzl", "cpp_proto_repositories")
cpp_proto_repositories()
