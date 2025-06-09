package Hamming;
use strict;
use warnings;
use feature qw(signatures);
use Exporter 'import';
our @EXPORT_OK = qw(hamming_distance);

sub hamming_distance ($strand1, $strand2) {
  die "left and right strands must be of equal length" if (length($strand1) != length($strand2));
  my $sum = 0;
  foreach my $i (0..(length($strand1)-1)) {
    $sum++ if(substr($strand1, $i, 1) ne substr($strand2, $i, 1));
  }
  return $sum;
}

1;


# vim:ts=2:sw=2:expandtab
