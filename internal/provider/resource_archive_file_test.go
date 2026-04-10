// Copyright IBM Corp. 2017, 2026
// SPDX-License-Identifier: MPL-2.0

package archive

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	r "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestResource_UpgradeFromVersion2_2_0_ContentConfig(t *testing.T) {
	td := t.TempDir()

	f := filepath.Join(td, "zip_file_acc_test_upgrade_content_config.zip")

	var fileSize string

	r.ParallelTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				ExternalProviders: map[string]r.ExternalProvider{
					"archive": {
						VersionConstraint: "2.2.0",
						Source:            "hashicorp/archive",
					},
				},
				Config: testAccArchiveFileResourceContentConfig("zip", f),
				Check: r.ComposeTestCheckFunc(
					testAccArchiveFileSize(f, &fileSize),
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_size", &fileSize),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_base64sha256", "P7VckxoEiUO411WN3nwuS/yOBL4zsbVWkQU9E1I5H6c=",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_md5", "ea35f0444ea9a3d5641d8760bc2815cc",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_sha", "019c79c4dc14dbe1edb3e467b2de6a6aad148717",
					),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccArchiveFileResourceContentConfig("zip", f),
				PlanOnly:                 true,
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccArchiveFileResourceContentConfig("zip", f),
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_size", &fileSize),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_base64sha256", "P7VckxoEiUO411WN3nwuS/yOBL4zsbVWkQU9E1I5H6c=",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_md5", "ea35f0444ea9a3d5641d8760bc2815cc",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_sha", "019c79c4dc14dbe1edb3e467b2de6a6aad148717",
					),
				),
			},
		},
	})
}

func TestResource_UpgradeFromVersion2_2_0_FileConfig(t *testing.T) {
	td := t.TempDir()

	f := filepath.Join(td, "zip_file_acc_test_upgrade_file_config.zip")

	var fileSize string

	r.ParallelTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				ExternalProviders: map[string]r.ExternalProvider{
					"archive": {
						VersionConstraint: "2.2.0",
						Source:            "hashicorp/archive",
					},
				},
				Config: testAccArchiveFileResourceFileConfig("zip", f),
				Check: r.ComposeTestCheckFunc(
					testAccArchiveFileSize(f, &fileSize),
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_size", &fileSize),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_base64sha256", "UTE4f5cWfaR6p0HfOrLILxgvF8UUwiJTjTRwjQTgdWs=",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_md5", "59fbc9e62af3cbc2f588f97498240dae",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_sha", "ce4ee1450ab93ac86e11446649e44cea907b6568",
					),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccArchiveFileResourceFileConfig("zip", f),
				PlanOnly:                 true,
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccArchiveFileResourceFileConfig("zip", f),
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_size", &fileSize),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_base64sha256", "UTE4f5cWfaR6p0HfOrLILxgvF8UUwiJTjTRwjQTgdWs=",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_md5", "59fbc9e62af3cbc2f588f97498240dae",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_sha", "ce4ee1450ab93ac86e11446649e44cea907b6568",
					),
				),
			},
		},
	})
}

func TestResource_UpgradeFromVersion2_2_0_DirConfig(t *testing.T) {
	td := t.TempDir()

	f := filepath.Join(td, "zip_file_acc_test_upgrade_dir_config.zip")

	var fileSize string

	r.ParallelTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				ExternalProviders: map[string]r.ExternalProvider{
					"archive": {
						VersionConstraint: "2.2.0",
						Source:            "hashicorp/archive",
					},
				},
				Config: testAccArchiveFileResourceDirConfig("zip", f),
				Check: r.ComposeTestCheckFunc(
					testAccArchiveFileSize(f, &fileSize),
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_size", &fileSize),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_base64sha256", "ydB8wtq8nK9vQ77VH6YTwoHmyljK46jW+uIJSwCzNpo=",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_md5", "b73f64a383716070aa4a29563b8b14d4",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_sha", "76d20a402eefd1cfbdc47886abd4e0909616c191",
					),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccArchiveFileResourceDirConfig("zip", f),
				PlanOnly:                 true,
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccArchiveFileResourceDirConfig("zip", f),
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_size", &fileSize),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_base64sha256", "ydB8wtq8nK9vQ77VH6YTwoHmyljK46jW+uIJSwCzNpo=",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_md5", "b73f64a383716070aa4a29563b8b14d4",
					),
					r.TestCheckResourceAttr(
						"archive_file.foo", "output_sha", "76d20a402eefd1cfbdc47886abd4e0909616c191",
					),
				),
			},
		},
	})
}

func TestResource_UpgradeFromVersion2_2_0_DirExcludesConfig(t *testing.T) {
	td := t.TempDir()

	f := filepath.Join(td, "zip_file_acc_test_upgrade_dir_excludes.zip")

	var fileSize, outputSha string

	r.ParallelTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				ExternalProviders: map[string]r.ExternalProvider{
					"archive": {
						VersionConstraint: "2.2.0",
						Source:            "hashicorp/archive",
					},
				},
				Config: testAccArchiveFileResourceDirExcludesConfig("zip", f),
				Check: r.ComposeTestCheckFunc(
					testAccArchiveFileSize(f, &fileSize),
					testExtractResourceAttr("archive_file.foo", "output_sha", &outputSha),
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_size", &fileSize),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccArchiveFileResourceDirExcludesConfig("zip", f),
				PlanOnly:                 true,
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccArchiveFileResourceDirExcludesConfig("zip", f),
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_size", &fileSize),
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_sha", &outputSha),
				),
			},
		},
	})
}

func TestResource_UpgradeFromVersion2_2_0_SourceConfig(t *testing.T) {
	td := t.TempDir()

	f := filepath.Join(td, "zip_file_acc_test_upgrade_source.zip")

	var fileSize, outputSha string

	r.ParallelTest(t, r.TestCase{
		Steps: []r.TestStep{
			{
				ExternalProviders: map[string]r.ExternalProvider{
					"archive": {
						VersionConstraint: "2.2.0",
						Source:            "hashicorp/archive",
					},
				},
				Config: testAccArchiveFileResourceMultiSourceConfig("zip", f),
				Check: r.ComposeTestCheckFunc(
					testAccArchiveFileSize(f, &fileSize),
					testExtractResourceAttr("archive_file.foo", "output_sha", &outputSha),
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_size", &fileSize),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccArchiveFileResourceMultiSourceConfig("zip", f),
				PlanOnly:                 true,
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccArchiveFileResourceMultiSourceConfig("zip", f),
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_size", &fileSize),
					r.TestCheckResourceAttrPtr("archive_file.foo", "output_sha", &outputSha),
				),
			},
		},
	})
}

func TestResource_SourceConfigMissing(t *testing.T) {
	r.ParallelTest(t, r.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),
		Steps: []r.TestStep{
			{
				Config:      testResourceSourceConfigMissing("zip"),
				ExpectError: regexp.MustCompile(`.*At least one of these attributes must be configured:\n\[source,source_content_filename,source_file,source_dir]`),
			},
		},
	})
}

func TestResource_SourceConfigConflicting(t *testing.T) {
	r.ParallelTest(t, r.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),
		Steps: []r.TestStep{
			{
				Config:      testResourceSourceConfigConflicting("zip"),
				ExpectError: regexp.MustCompile(`.*Attribute "source_dir" cannot be specified when "source" is specified`),
			},
		},
	})
}

func alterFileContents(content, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(fmt.Sprintf("error creating file: %s", err))
	}

	defer f.Close()

	_, err = f.Write([]byte(content))
	if err != nil {
		panic(fmt.Sprintf("error writing file: %s", err))
	}
}

func testAccArchiveFileResourceContentConfig(format, outputPath string) string {
	return fmt.Sprintf(`
resource "archive_file" "foo" {
  type                    = "%s"
  source_content          = "This is some content"
  source_content_filename = "content.txt"
  output_path             = "%s"
}
`, format, filepath.ToSlash(outputPath))
}

func testAccArchiveFileResourceFileConfig(format, outputPath string) string {
	return fmt.Sprintf(`
resource "archive_file" "foo" {
  type             = "%s"
  source_file      = "test-fixtures/test-dir/test-file.txt"
  output_path      = "%s"
  output_file_mode = "0666"
}
`, format, filepath.ToSlash(outputPath))
}

func testAccArchiveFileResourceFileSourceFileConfig(format, sourceFile, outputPath string) string {
	return fmt.Sprintf(`
resource "archive_file" "foo" {
  type             = "%s"
  source_file      = "%s"
  output_path      = "%s"
  output_file_mode = "0666"
}
`, format,
		filepath.ToSlash(sourceFile),
		filepath.ToSlash(outputPath))
}

func testAccArchiveFileResourceDirConfig(format, outputPath string) string {
	return fmt.Sprintf(`
resource "archive_file" "foo" {
  type             = "%s"
  source_dir       = "test-fixtures/test-dir/test-dir1"
  output_path      = "%s"
  output_file_mode = "0666"
}
`, format, filepath.ToSlash(outputPath))
}

func testAccArchiveFileResourceDirExcludesConfig(format, outputPath string) string {
	return fmt.Sprintf(`
resource "archive_file" "foo" {
  type        = "%s"
  source_dir  = "test-fixtures/test-dir"
  excludes    = ["test-fixtures/test-dir/file2.txt"]
  output_path = "%s"
}
`, format, filepath.ToSlash(outputPath))
}

func testAccArchiveFileResourceDirExcludesGlobConfig(format, outputPath string) string {
	return fmt.Sprintf(`
resource "archive_file" "foo" {
  type        = "%s"
  source_dir  = "test-fixtures/test-dir"
  excludes    = ["test-fixtures/test-dir/file2.txt", "**/file[2-3].txt"]
  output_path = "%s"
}
`, format, filepath.ToSlash(outputPath))
}

func testAccArchiveFileResourceMultiSourceConfig(format, outputPath string) string {
	return fmt.Sprintf(`
resource "archive_file" "foo" {
  type = "%s"
  source {
    filename = "content_1.txt"
    content = "This is the content for content_1.txt"
  }
  source {
    filename = "content_2.txt"
    content = "This is the content for content_2.txt"
  }
  output_path = "%s"
}
`, format, filepath.ToSlash(outputPath))
}

func testResourceSourceConfigMissing(format string) string {
	return fmt.Sprintf(`
resource "archive_file" "foo" {
  type                    = "%s"
  output_path             = "path"
}
`, format)
}

func testResourceSourceConfigConflicting(format string) string {
	return fmt.Sprintf(`
resource "archive_file" "foo" {
  type                    = "%s"
  source {
    filename = "content_1.txt"
    content = "This is the content for content_1.txt"
  }
  source_dir  = "test-fixtures/test-dir"
  output_path             = "path"
}
`, format)
}

// TestResource_OutputDeleteAndSourceChange_Matrix verifies the archive_file resource
// works correctly in CI workflows where the output file may be deleted between runs.
// It covers all 4 combinations of {file deleted, file present} x {source unchanged,
// source changed} and ensures a downstream consumer (terraform_data) only churns when
// the archive inputs change.
func TestResource_OutputDeleteAndSourceChange_Matrix(t *testing.T) {
	td := t.TempDir()
	outputPath := filepath.Join(td, "matrix_test.zip")

	var originalMd5, originalSha string
	var consumerId string

	configFunc := func(content string) string {
		return fmt.Sprintf(`
resource "archive_file" "this" {
  type                    = "zip"
  source_content          = "%s"
  source_content_filename = "content.txt"
  output_path             = "%s"
}

resource "terraform_data" "consumer" {
  input = {
    name       = "obj-${archive_file.this.output_md5}.zip"
    source     = archive_file.this.output_path
    source_md5 = archive_file.this.output_md5
  }
}
`, content, filepath.ToSlash(outputPath))
	}

	deleteOutputFile := func() {
		_ = os.Remove(outputPath)
	}

	r.Test(t, r.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),
		Steps: []r.TestStep{
			// Step 1: Create with content="original". Record hashes and consumer id.
			{
				Config: configFunc("original"),
				Check: r.ComposeTestCheckFunc(
					testExtractResourceAttr("archive_file.this", "output_md5", &originalMd5),
					testExtractResourceAttr("archive_file.this", "output_sha", &originalSha),
					testExtractResourceAttr("terraform_data.consumer", "id", &consumerId),
				),
			},
			// Step 2: File present + source unchanged -> PlanOnly, expect empty plan.
			{
				Config:   configFunc("original"),
				PlanOnly: true,
			},
			// Step 3: File DELETED + source unchanged -> PlanOnly, expect non-empty plan.
			{
				PreConfig:          deleteOutputFile,
				Config:             configFunc("original"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// Step 4: File DELETED + source unchanged -> apply.
			//   Archive re-created with same hashes. Consumer should NOT churn.
			{
				PreConfig: deleteOutputFile,
				Config:    configFunc("original"),
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttrPtr("archive_file.this", "output_md5", &originalMd5),
					r.TestCheckResourceAttrPtr("archive_file.this", "output_sha", &originalSha),
					r.TestCheckResourceAttrPtr("terraform_data.consumer", "id", &consumerId),
					r.TestCheckResourceAttrPtr("terraform_data.consumer", "output.source_md5", &originalMd5),
				),
			},
			// Step 5: File present + source unchanged -> no changes at all.
			{
				Config: configFunc("original"),
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttrPtr("archive_file.this", "output_md5", &originalMd5),
					r.TestCheckResourceAttrPtr("archive_file.this", "output_sha", &originalSha),
					r.TestCheckResourceAttrPtr("terraform_data.consumer", "id", &consumerId),
					r.TestCheckResourceAttrPtr("terraform_data.consumer", "output.source_md5", &originalMd5),
				),
			},
			// Step 6: File present + source CHANGED -> new hashes, consumer churns.
			{
				Config: configFunc("modified"),
				Check: r.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["archive_file.this"]
						if !ok {
							return fmt.Errorf("archive_file.this not found in state")
						}
						newMd5 := rs.Primary.Attributes["output_md5"]
						if newMd5 == originalMd5 {
							return fmt.Errorf("expected output_md5 to change after source modification, but it stayed %s", originalMd5)
						}
						return nil
					},
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["terraform_data.consumer"]
						if !ok {
							return fmt.Errorf("terraform_data.consumer not found in state")
						}
						consumerMd5 := rs.Primary.Attributes["output.source_md5"]
						if consumerMd5 == originalMd5 {
							return fmt.Errorf("expected consumer output.source_md5 to change after source modification, but it stayed %s", originalMd5)
						}
						return nil
					},
				),
			},
			// Step 7: File DELETED + source unchanged (still "modified") -> re-create,
			//   same hashes as step 6. Consumer should NOT churn.
			{
				PreConfig: deleteOutputFile,
				Config:    configFunc("modified"),
				Check: r.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["archive_file.this"]
						if !ok {
							return fmt.Errorf("archive_file.this not found in state")
						}
						newMd5 := rs.Primary.Attributes["output_md5"]
						if newMd5 == originalMd5 {
							return fmt.Errorf("expected output_md5 to differ from original %s", originalMd5)
						}
						return nil
					},
					r.TestCheckResourceAttrPtr("terraform_data.consumer", "id", &consumerId),
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["terraform_data.consumer"]
						if !ok {
							return fmt.Errorf("terraform_data.consumer not found in state")
						}
						consumerMd5 := rs.Primary.Attributes["output.source_md5"]
						if consumerMd5 == originalMd5 {
							return fmt.Errorf("expected consumer output.source_md5 to differ from original %s", originalMd5)
						}
						return nil
					},
				),
			},
		},
	})
}

// TestResource_TamperedOutputFile verifies that when the output file is replaced
// with different contents between runs, Read detects the mismatch with the expected
// archive and removes the resource from state, triggering a re-create on apply.
func TestResource_TamperedOutputFile(t *testing.T) {
	td := t.TempDir()
	outputPath := filepath.Join(td, "tampered_test.zip")

	var originalMd5 string

	config := fmt.Sprintf(`
resource "archive_file" "this" {
  type                    = "zip"
  source_content          = "hello"
  source_content_filename = "content.txt"
  output_path             = "%s"
}

resource "terraform_data" "consumer" {
  input = archive_file.this.output_md5
}
`, filepath.ToSlash(outputPath))

	tamperOutputFile := func() {
		// Overwrite the output file with arbitrary different content.
		if err := os.WriteFile(outputPath, []byte("tampered data"), 0644); err != nil {
			t.Fatalf("failed to tamper output file: %s", err)
		}
	}

	r.Test(t, r.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),
		Steps: []r.TestStep{
			// Step 1: Create the archive normally. Record the md5.
			{
				Config: config,
				Check: r.ComposeTestCheckFunc(
					testExtractResourceAttr("archive_file.this", "output_md5", &originalMd5),
				),
			},
			// Step 2: Tamper with the output file. Read should detect the
			// checksum mismatch and remove from state, causing a non-empty
			// plan (re-create archive_file + update consumer).
			{
				PreConfig:          tamperOutputFile,
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// Step 3: Apply after tampering. The archive is re-created from
			// the unchanged sources, so it should have the original md5.
			{
				PreConfig: tamperOutputFile,
				Config:    config,
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttrPtr("archive_file.this", "output_md5", &originalMd5),
				),
			},
		},
	})
}
