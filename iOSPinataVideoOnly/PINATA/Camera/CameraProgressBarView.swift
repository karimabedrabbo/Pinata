//
//  CameraProgressBarView.swift
//  PINATA
//
//  Created by Karim Abedrabbo on 11/19/19.
//  Copyright © 2019 Karim Abedrabbo. All rights reserved.
//

import Foundation
import UIKit

public class CameraProgressBarView: UIView {
  
  
     var progressViewBars: [UIView] = []
     var colorRotation: [UIColor] = [.green, .blue, .red,  .magenta, .yellow, .cyan, .brown, .purple, .orange]
     
     var alphaIsIncreasing: Bool = false
     var readyToAddBar: Bool = false
     var lastCompletionRatio: CGFloat = 0.0
     var completionRatios: [CGFloat] = []
     var activeBarWidthConstraint: NSLayoutConstraint?
  
  var progressView: UIView = {
         let progress = UIView()
         progress.backgroundColor = UIColor.white.withAlphaComponent(0.3)
         progress.contentMode = .left
         progress.layer.cornerRadius = 15.0
         progress.translatesAutoresizingMaskIntoConstraints = false
         progress.clipsToBounds = true
         return progress
     }()
     
     var minimumTimeView: UIView = {
         let minimumTime = UIView()
         minimumTime.backgroundColor = UIColor.white.withAlphaComponent(0.6)
         minimumTime.translatesAutoresizingMaskIntoConstraints = false
         return minimumTime
     }()
}
