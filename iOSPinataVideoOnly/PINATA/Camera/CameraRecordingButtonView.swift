//
//  CameraRecordingButtonView.swift
//  PINATA
//
//  Created by Karim Abedrabbo on 11/19/19.
//  Copyright © 2019 Karim Abedrabbo. All rights reserved.
//

import Foundation
import UIKit
public class CameraRecordingButtonView: UIView {
  
    var blurredRecordView: UIView = {
        let blurEffect = UIBlurEffect(style: UIBlurEffect.Style.regular)

        let blurView = UIVisualEffectView(effect: blurEffect)


        blurView.backgroundColor = UIColor.clear
        blurView.translatesAutoresizingMaskIntoConstraints = false
        return blurView
    }()
  
    var innerShapeLayer: CAShapeLayer =  {
        let innerShape = CAShapeLayer()
        innerShape.backgroundColor = UIColor.clear.cgColor
        innerShape.fillColor = UIColor.clear.cgColor
        return innerShape
    }()

    var outerShapeLayer: CAShapeLayer =  {
        let outerShape = CAShapeLayer()
        outerShape.lineWidth = 4.0
        outerShape.backgroundColor = UIColor.clear.cgColor
        outerShape.strokeColor = UIColor.white.cgColor
        outerShape.fillColor = UIColor.clear.cgColor
        return outerShape
    }()
  
  
  public override init(frame: CGRect) {
      super.init(frame: .zero)
      self.layer.addSublayer(outerShapeLayer)
      self.addSubview(blurredRecordView)
      blurredRecordView.layer.mask = innerShapeLayer
  
    NSLayoutConstraint.activate([
      
    self.blurredRecordView.centerXAnchor.constraint(equalTo: self.centerXAnchor),
    self.blurredRecordView.centerYAnchor.constraint(equalTo: self.centerYAnchor),
    self.blurredRecordView.heightAnchor.constraint(equalTo: self.heightAnchor, constant: 0),
    self.blurredRecordView.widthAnchor.constraint(equalTo: self.widthAnchor, constant: 0),
    
    ])
          
  }
  
  required init?(coder: NSCoder) {
    fatalError("init(coder:) has not been implemented")
  }
  
  override public func layoutSubviews() {
      super.layoutSubviews()
      self.updateOuter()
      self.updateInner()
      
  }
  
  
  func updateOuter() {
         let recordViewCenter = CGPoint(x: self.bounds.width / 2, y: self.bounds.height / 2)
         let outerPath = UIBezierPath(arcCenter: recordViewCenter, radius: self.bounds.width / 3, startAngle: 0.0, endAngle: 2.0 * .pi, clockwise: false)
         outerShapeLayer.path = outerPath.cgPath
         
     }
     
     func updateInner() {
         let blurredRecordViewCenter = CGPoint(x: blurredRecordView.bounds.width / 2, y: blurredRecordView.bounds.height / 2)
         let innerPath = UIBezierPath(arcCenter: blurredRecordViewCenter, radius: blurredRecordView.bounds.width / 4, startAngle: 0.0, endAngle: 2.0 * .pi, clockwise: false)
         innerShapeLayer.path = innerPath.cgPath
     }
  
}
